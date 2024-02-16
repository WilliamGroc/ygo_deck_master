import { ActionFunctionArgs, LoaderFunctionArgs, json, redirect } from "@remix-run/node";
import { Form, Link } from "@remix-run/react";
import { axiosInstance } from "~/utils/axios.server";
import { commitSession, getSession } from "~/utils/session.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const session = await getSession(
    request.headers.get("Cookie")
  );

  if (session.has("userId")) {
    return redirect("/");
  }

  const data = { error: session.get("error") };

  return json(data, {
    headers: {
      "Set-Cookie": await commitSession(session),
    },
  });
}

export async function action({ request }: ActionFunctionArgs) {
  const formData = await request.formData();

  const email = formData.get("email");
  const password = formData.get("password");

  const session = await getSession(
    request.headers.get("Cookie")
  );

  try {
    const { data } = await axiosInstance.post("/users/login", { email, password });
    console.log(data);
    session.set("token", data.token);

    return redirect("/",
      {
        headers: {
          "Set-Cookie": await commitSession(session),
        }
      });
  } catch (error) {
    session.flash("error", "Invalid username/password");

    return redirect("/login", {
      headers: {
        "Set-Cookie": await commitSession(session),
      },
    });
  }
}

export default function LoginPage() {
  return (
    <div>
      <h1>Login</h1>
      <Form method="post">
        <label>
          Email
          <input type="email" name="email" required />
        </label>
        <label>
          Password
          <input type="password" name="password" required />
        </label>
        <button type="submit">Login</button>
        <Link to="/register" className="btn btn-secondary">
          Register
        </Link>
      </Form>
    </div>
  );
}