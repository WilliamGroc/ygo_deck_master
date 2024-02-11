import { ActionFunctionArgs } from "@remix-run/node";
import { Form, Link } from "@remix-run/react";
import axios from "axios";

export async function action({ request }: ActionFunctionArgs) {
  const formData = await request.formData();

  const email = formData.get("email");
  const password = formData.get("password");

  try {
    const { data } = await axios.post("http://localhost:8080/users/login", { email, password });
    console.log({ data });
    return { token: data.token };
  } catch (error) {
    return new Response("Failed to login", { status: 400 });
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