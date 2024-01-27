"use client";

import type { LoaderFunction, MetaFunction } from "@remix-run/node";
import { json, useLoaderData, useNavigate, useSearchParams } from "@remix-run/react";
import axios from "axios";
import { CardTile } from "~/components/cardTile";
import { Card } from "~/models/card.model";
import styles from "../styles/_index.module.css";
import { useState } from "react";

export const loader: LoaderFunction = async ({ request }) => {
  const search = new URL(request.url).searchParams.get('search');
  const { data } = await axios.get<Card[]>("http://localhost:8080/cards", {params: {search}});
  return json({ data });
}

export const meta: MetaFunction = () => {
  return [
    { title: "Card list" },
    { name: "description", content: "List of card yugioh" },
  ];
};

export default function Index() {
  const { data } = useLoaderData<{ data: Card[] }>();

  const [searchParams] = useSearchParams({ search: '' });
  const navigate = useNavigate();

  const [search, setSearch] = useState(searchParams.get('search') || '');

  const handleSearch = () => {
    navigate(`?search=${search}`);
  }

  return (
    <div>
      <div className="title">Card list</div>
      <div className="flex">
        <input type="text" placeholder="Search" onInput={e => setSearch(e.currentTarget.value)} />
        <div>
          <button onClick={handleSearch}>Search</button>
        </div>
      </div>
      <div className={styles['card-list-container']}>
        {data.map((card) => (
          <CardTile key={card.id} card={card} />
        ))}
      </div>
    </div>
  );
}
