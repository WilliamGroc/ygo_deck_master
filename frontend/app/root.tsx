import { cssBundleHref } from "@remix-run/css-bundle";
import type { LinksFunction, LoaderFunctionArgs } from "@remix-run/node";
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  redirect,
  useFetcher,
  useLoaderData,
} from "@remix-run/react";

import tailwind from "./tailwind.css";
import styles from "./styles.css";
import { Navbar } from "./components/navbar";
import { langCookie } from "./utils/cookie.server";
import { Lang } from "./const/lang";
import { getSession } from "./utils/session.server";
import axios from "axios";

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: tailwind },
  { rel: "stylesheet", href: styles },
  ...(cssBundleHref ? [{ rel: "stylesheet", href: cssBundleHref }] : []),
];

export async function loader({ request }: LoaderFunctionArgs) {
  const cookie = request.headers.get('Cookie');
  const lang = await langCookie.parse(cookie) || 'en';

  const session = await getSession(
    request.headers.get("Cookie")
  );

  if (session.has("token")) {
    axios.defaults.headers.common['Authorization'] = 'Bearer ' + session.get("token");
  }

  return { lang, isAuthenticatied: session.has("token") };
}

export async function action({ request }: LoaderFunctionArgs) {
  const data = await request.formData();
  return redirect(data.get('redirectTo') as string, { headers: { 'Set-Cookie': await langCookie.serialize(data.get('lang')) } });
}

export default function App() {
  const { lang, isAuthenticatied } = useLoaderData<ReturnType<typeof loader>>();
  const fetcher = useFetcher();

  const setLang = async (lang: Lang) => {
    console.log('setLang', lang);
    fetcher.submit({ lang, redirectTo: location.href }, { method: 'POST' });
  }

  return (
    <html lang={lang}>
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        <Navbar setLang={setLang} currentLang={lang} isAuthenticatied={isAuthenticatied} />
        <div className="p-3">
          <Outlet />
        </div>
        <ScrollRestoration />
        <Scripts />
        <LiveReload />
      </body>
    </html>
  );
}
