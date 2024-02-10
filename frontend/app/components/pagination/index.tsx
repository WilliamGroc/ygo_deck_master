
type Props = {
  total: number,
  perPage: number,
  currentPage: number,
  onPageChange: (page: number) => void
}

export function Pagination({ total, perPage, currentPage, onPageChange}: Props) {
  return (
    <div className="flex items-center">
      <button className="round" onClick={() => onPageChange(currentPage - 1)} disabled={currentPage === 1}>Previous</button>
      <span className="ml-4 mr-4">{currentPage} / {Math.ceil(total/perPage)}</span>
      <button className="round" onClick={() => onPageChange(currentPage + 1)} disabled={currentPage === Math.ceil(total / perPage)}>Next</button>
    </div>
  );
}