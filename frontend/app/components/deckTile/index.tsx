import { Deck } from "~/models/deck.model";
import styles from "./styles.module.css";

type Props = {
  deck: Deck;
};

export default function DeckTile({ deck }: Props) {
  return (
    <div className={styles['deck-tile']}>
      <div className={styles['deck-tile__info-container']}>
        <div className={styles['deck-tile__name']}>
          {deck.name}
        </div>
      </div>
    </div>
  );
}