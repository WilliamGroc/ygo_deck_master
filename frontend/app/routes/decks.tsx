import { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { Deck } from "~/models/deck.model";
import { axiosInstance } from "~/utils/axios.server";
import { getSession } from "~/utils/session.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const session = await getSession(
    request.headers.get("Cookie")
  );

  const userId = session.get("userId");

  const { data } = await axiosInstance.get<Deck[]>('/decks', {
    params: {
      userId
    }
  });

  return { decks: data };
}

export default function DecksPage() {
  const { decks } = useLoaderData<ReturnType<typeof loader>>();

  return (
    <div>
      <h1>Decks</h1>
      <ul>
        {decks.map(deck => (
          <li key={deck.id}>{deck.name}</li>
        ))}
      </ul>
    </div>
  );
}