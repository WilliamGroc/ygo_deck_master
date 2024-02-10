import { Card } from "~/models/card.model"
import styles from "./styles.module.css";
import { Lang } from "~/const/lang";
import { useMemo } from "react";

type Props = {
  card: Card,
  lang: Lang
}

export function CardTile({ card, lang }: Props) {
  const translatedCard = useMemo(() => {
    if (card[lang]?.name)
      return card[lang];
    else
      return card.en;
  }, [card, lang]);
  
  return <div className={styles['card-tile']}>
    <div className={styles['card-tile__image']}>
      <img src={`http://localhost:8080/cards/${card.id}/image`} width={135} />
    </div>
    <div className={styles['card-tile__info-container']}>
      <div>
        <div className={styles['card-tile__name']}>
          {translatedCard.name}
        </div>
        <div className="capitalize">
          {[card.race, card.type, card.frameType].join(' / ')}
        </div>
      </div>
      {
        !['trap', 'spell'].includes(card.frameType) &&
        <div className="flex flex-col">
          <span>Atk: {card.atk}</span>
          <span>Def: {card.def}</span>
        </div>
      }
    </div>
  </div>
}