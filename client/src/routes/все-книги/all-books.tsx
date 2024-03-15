import { useNavigate } from "@solidjs/router";
import useBooks from "../../hooks/useBooks";

const AllBooks = () => {
  const { books } = useBooks();
  const navigate = useNavigate();

  const handleClick = () => {
    navigate("/добавить-книгу")
  }

  return (
    <div class="container mx-auto pt-5 w-[95%]">
      <div class="flex center flex-col">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-2xl font-bold">Все книги</h2>
          <button onClick={handleClick}>Добавить книгу</button>
        </div>
        <div class="overflow-x-auto rounded">
          <table class="table-auto border-collapse border">
            <thead>
              <tr>
                <th class="border p-3">Название</th>
                <th class="border p-3">Автор</th>
                <th class="border p-3">Год издания</th>
              </tr>
            </thead>
            <tbody>
              {books.map((book, index) => (
                <tr>
                  <td class="border p-3">{book.title}</td>
                  <td class="border p-3">{book.author}</td>
                  <td class="border p-3">{book.year}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default AllBooks;
