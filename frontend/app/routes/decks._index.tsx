import { LoaderFunctionArgs } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";
import DeckTile from "~/components/deckTile";
import { Deck } from "~/models/deck.model";
import { axiosInstance } from "~/utils/axios.server";
import { getSession } from "~/utils/session.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const session = await getSession(
    request.headers.get("Cookie")
  );

  const userId = session.get("userId");

  try {
    const { data: { data, total } } = await axiosInstance.get<{ data: Deck[], total: number, filters: any }>('/decks', {
      params: {
        userId
      }
    });

    return { decks: data, total };
  }
  catch (error) {
    console.error(error);
    throw new Error("Failed to load decks");
  }
}

export default function DecksPage() {
  const { decks } = useLoaderData<ReturnType<typeof loader>>();

  return (
    <div>
      <h1>Decks</h1>
      <div className="flex flex-wrap justify-between">
        {decks.map(deck => (
          <Link to={`/decks/${deck.id}`} style={{ width: '32%' }} className="mb-2" key={deck.id}>
            <DeckTile deck={deck} />
          </Link>
        ))}
      </div>
    </div>
  );
}