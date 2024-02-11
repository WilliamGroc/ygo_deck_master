import { ActionFunctionArgs } from "@remix-run/node";
import { Form, Link } from "@remix-run/react";

export async function action({ request }: ActionFunctionArgs) {
  const formData = await request.formData();

  const email = formData.get("email");
  const password = formData.get("password");

  const response = await fetch("http://localhost:8080/users/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  });

  if (response.ok) {
    return response;
  }

  return new Response("Failed to register", { status: 400 });
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