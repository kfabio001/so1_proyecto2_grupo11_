import {BrowserRouter, Route, Routes } from 'react-router-dom';

import Layout from './Layout';


import OneDose from '../pages/OneDose';
import TwoDose from '../pages/TwoDose';
import TodoData from '../pages/TodoData';
import InfoData from '../pages/InfoData';

function App() {
  return (
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path="/one-dose" element={<OneDose />} />
          <Route path="/two-dose" element={<TwoDose />} />
          <Route path="/all-data" element={<TodoData />} />
          <Route path="/info-data" element={<InfoData />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  );
}

export default App;
