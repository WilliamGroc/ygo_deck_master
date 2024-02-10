import { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import axios from "axios";
import { useEffect, useMemo, useState } from "react";
import { Lang } from "~/const/lang";
import { Card } from "~/models/card.model";
import { langCookie } from "~/utils/cookie.server";

export async function loader({ params, request }: LoaderFunctionArgs) {
  const cookie = request.headers.get('Cookie');
  const lang: Lang = await langCookie.parse(cookie) || 'en';

  const { data } = await axios.get<Card>(`http://localhost:8080/cards/${params.id}`);
  return { card: data, lang };
}


export default function CardPage() {
  const { card, lang } = useLoaderData<ReturnType<typeof loader>>();

  const [width, setWidth] = useState(100);

  useEffect(() => {
    setWidth(document.getElementsByTagName('body')[0].clientWidth * 0.2);
  }, []);

  const translatedCard = useMemo(() => {
    if (card[lang]?.name)
      return card[lang];
    else
      return card.en;
  }, [card, lang]);

  return (
    <div className="flex">
      <img src={`http://localhost:8080/cards/${card.id}/image/big`} width={width} />
      <div className="p-4 flex-1">
        <div className="font-bold text-xl">
          {translatedCard.name}
        </div>
        <div>
          {
            !['trap', 'spell'].includes(card.frameType) &&
            <div className="flex flex-col">
              <span>Atk: <b>{card.atk}</b></span>
              <span>Def: <b>{card.def}</b></span>
            </div>
          }
        </div>
        <div>
          {
            translatedCard.effectText
          }
        </div>
      </div>
    </div>
  )
}