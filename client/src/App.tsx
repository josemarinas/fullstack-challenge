import Pages from './types';
import { useState } from 'react';
import { HomePage, NotFoundPage } from './pages';
import { QueryClient, QueryClientProvider } from "react-query";
const queryClient = new QueryClient();
function App() {
  const [currentPage, setCurrentPage] = useState<Pages>(Pages.Home)
  const onSetPage = (page: Pages) => {
    setCurrentPage(page)
  }
  switch (currentPage) {
    case Pages.Home:

      return (
        <QueryClientProvider client={queryClient}>
          <HomePage onSetPage={onSetPage} />
        </QueryClientProvider>
      )
    default:
      return (
        <QueryClientProvider client={queryClient}>
          <NotFoundPage onSetPage={onSetPage} />
        </QueryClientProvider>
      )
  }
}

export default App;
