import { Card } from "~/models/card.model"
import styles from "./styles.module.css";

type Props = {
  card: Card
}

export function CardTile({ card }: Props) {
  return <div className={styles['card-tile']}>
    <div className={styles['card-tile__image']}>
      <img src={`http://localhost:8080/cards/${card.id}/image`} width={135} />
    </div>
    <div className={styles['card-tile__info-container']}>
      <div className={styles['card-tile__name']}>
        {card.name}
      </div>
      <div className="flex flex-col">
        <span>Atk: {card.atk}</span>
        <span>Def: {card.def}</span>
      </div>
    </div>
  </div>
}