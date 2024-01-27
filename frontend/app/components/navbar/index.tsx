import { Link } from "@remix-run/react";

export function Navbar() {
  return (
    <div className="flex flex-row justify-between items-center p-4 bg-gray-200">
      <div className="text-xl font-bold">Yugioh deck master</div>
      <div className="flex flex-row space-x-2">
        <Link to="/" className="btn btn-secondary">
          Card list
        </Link>
      </div>
    </div>
  );
}