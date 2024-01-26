import { Title } from "@solidjs/meta";
import { For } from "solid-js";
import { CardTile } from "~/components/cardTile";
import { useCards } from "~/services/card";

export default function Home() {
  const cardsService = useCards();

  return (
    <main>
      <Title>Card List</Title>
      <h1>Card list</h1>

      <ul>
        <For each={cardsService.data?.data}>
          {(card) => (
            <li>
              <a href={`/card/${card.id}`}><CardTile card={card} /></a>
            </li>
          )}
        </For>
      </ul>
    </main>
  );
}
