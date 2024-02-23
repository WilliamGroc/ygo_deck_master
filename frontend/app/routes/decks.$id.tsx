import { ActionFunctionArgs } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";

export function loader({ params }: ActionFunctionArgs) {
  return { deck: {} };
}

export default function DeckPage() {
  const { deck } = useLoaderData<ReturnType<typeof loader>>();

  return (
    <div>
      <Link to="/decks">Back to Decks</Link>
    </div>
  );
}