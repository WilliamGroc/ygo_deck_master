import { ActionFunctionArgs, LoaderFunctionArgs, json, redirect } from "@remix-run/node";
import { Form, Link } from "@remix-run/react";
import { axiosInstance } from "~/utils/axios.server";
import { commitSession, getSession } from "~/utils/session.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const session = await getSession(
    request.headers.get("Cookie")
  );

  if (session.has("token")) {
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
  const session = await getSession(
    request.headers.get("Cookie")
  );
  const formData = await request.formData();

  const email = formData.get("email");
  const password = formData.get("password");

  try {
    const { data } = await axiosInstance.post("/users/register", { email, password });

    session.set("token", data.token);

    return redirect("/", {
      headers: {
        "Set-Cookie": await commitSession(session),
      },
    });
  }
  catch (error) {
    session.flash("error", "Registration failed. Please try again.");

    return redirect("/register", {
      headers: {
        "Set-Cookie": await commitSession(session),
      },
    });

  }
}

export default function RegisterPage() {
  return (
    <div>
      <h1>Register</h1>
      <Form method="post">
        <label>
          Email
          <input type="email" name="email" required />
        </label>
        <label>
          Password
          <input type="password" name="password" required />
        </label>
        <button type="submit">Register</button>
      </Form>
      <Link to="/login" className="btn btn-secondary">
        You have an account ? Login
      </Link>
    </div>
  );
}