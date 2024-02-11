"use client";

import type { LoaderFunctionArgs, MetaFunction } from "@remix-run/node";
import { Link, useLoaderData, useNavigate, useSearchParams } from "@remix-run/react";
import { CardTile } from "~/components/cardTile";
import { Card } from "~/models/card.model";
import styles from "../styles/_index.module.css";
import { useState } from "react";
import { langCookie } from "~/utils/cookie.server";
import { Pagination } from "~/components/pagination";
import { axiosInstance } from "~/utils/axios.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const searchParams = new URL(request.url).searchParams;

  const search = searchParams.get('search');
  const page = Number(searchParams.get('page') || 1);
  const type = searchParams.get('type');
  const attribute = searchParams.get('attribute');
  const level = searchParams.get('level') || '';

  const cookie = request.headers.get('Cookie');
  const lang = await langCookie.parse(cookie) || 'en';

  const { data: { data, total, filters } } = await axiosInstance.get<{ data: Card[], total: number, filters: any }>("/cards", {
    params: {
      search,
      page,
      type,
      level,
      attribute
    }
  });
  return { data, lang, total, filters };
}

export const meta: MetaFunction = () => {
  return [
    { title: "Card list" },
    { name: "description", content: "List of card yugioh" },
  ];
};

export default function Index() {
  const { data, lang, total, filters } = useLoaderData<ReturnType<typeof loader>>();

  const [searchParams] = useSearchParams({ search: '', page: '', type: '', level: '', attribute: '' });
  const navigate = useNavigate();

  const [search, setSearch] = useState(searchParams.get('search') || '');

  const handlerFilterChange = (filterName: string) => {
    return (value: number | string) => {
      const queryParams = new URLSearchParams();
      queryParams.set('search', search);
      queryParams.set('page', String(1));
      queryParams.set('type', searchParams.get('type') || '');
      queryParams.set('level', searchParams.get('level') || '');
      queryParams.set('attribute', searchParams.get('attribute') || '');
      queryParams.set(filterName, value.toString());
      navigate(`?${queryParams.toString()}`);
    }
  }

  return (
    <div>
      <div className="title">Card list</div>
      <div className="flex items-stretch mb-4">
        <div className="flex w-1/3 h-full mt-5">
          <input type="text" placeholder="Search" onInput={e => setSearch(e.currentTarget.value)} onKeyUp={(e) => { if (e.key === "Enter") handlerFilterChange('search')(search); }} />
          <button className="ml-4" onClick={() => handlerFilterChange('search')(search)}>Search</button>
        </div>
        <div className="ml-4 flex w-full">
          <div className="w-1/4">
            <div>
              <label htmlFor="type">
                Types
              </label>
              <select className="ml-4 capitalize" value={searchParams.get('type') || ''} onChange={(e) => handlerFilterChange('type')(e.target.value)}>
                <option value="">All</option>
                {filters?.types.sort().map((type: string) => (
                  <option key={type} value={type} className="capitalize">{type}</option>
                ))}
              </select>
            </div>
            <div>
              <label htmlFor="level">
                Level
              </label>
              <select className="ml-4" value={searchParams.get('level') || ''} onChange={(e) => handlerFilterChange('level')(e.target.value)}>
                <option value="">All</option>
                {Array.from(Array(14).keys()).map((level: number) => (
                  <option key={level} value={level}>{level}</option>
                ))}
              </select>
            </div>
          </div>
          <div className="w-1/4 ml-4">
            <div>
              <label htmlFor="attribute">
                Attributes
              </label>
              <select className="ml-4 capitalize" value={searchParams.get('attribute') || ''} onChange={(e) => handlerFilterChange('attribute')(e.target.value)}>
                <option value="">All</option>
                {filters?.attributes.filter(Boolean).sort().map((type: string) => (
                  <option key={type} value={type} className="capitalize">{type}</option>
                ))}
              </select>
            </div>
          </div>
        </div>
      </div>
      <div className="mb-5">
        <Pagination total={total} perPage={20} currentPage={Number(searchParams.get('page')) || 1} onPageChange={handlerFilterChange('page')} />
      </div>
      <div className={styles['card-list-container']}>
        {data?.map((card) => (
          <Link to={`/cards/${card.id}`} key={card.id}>
            <CardTile key={card.id} card={card} lang={lang} />
          </Link>
        ))}
      </div>
    </div>
  );
}
