import { Card } from "~/models/card.model";

type Props = {
  card: Card
}

export function CardTile(props: Props) {
  const {
    card
  } = props;

  if (!card) {
    return null;
  }

  return <div>
    <img src={`http://localhost:8080/cards/${card.id}/image`} alt={card.name} />
    <h2>{card.name}</h2>
  </div>
}