import { Link } from "@remix-run/react";
import { Lang, LangEnum } from "~/const/lang";

type Props = {
  setLang: (lang: Lang) => void,
  currentLang: Lang
}

export function Navbar({ setLang, currentLang }: Props) {
  return (
    <div className="flex flex-row justify-between items-center p-4 bg-gray-200">
      <div className="flex items-center">
        <div className="text-xl font-bold">Yugioh deck master</div>
        <div className="flex flex-row space-x-2 ml-6">
          <Link to="/" className="btn btn-secondary">
            Card list
          </Link>
        </div>
      </div>
      <div>
        <div>
          <select className="p-2 uppercase" onChange={(e) => setLang(e.target.value as Lang)} value={currentLang}>
            {Object.values(LangEnum).map(lang => (
              <option key={lang} value={lang} className="uppercase">{lang}</option>
            ))}
          </select>
        </div>
      </div>
    </div>
  );
}