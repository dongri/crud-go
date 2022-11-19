import {BrowserRouter, Routes, Route} from "react-router-dom";
import { HelmetProvider } from 'react-helmet-async';

import { Header } from "./components/Header";
import { NotFound } from './components/NotFound';

import { Home } from './pages/Home';
import { List as PostsList } from './pages/posts/List';
import { New as PostsNew } from './pages/posts/New';
import { Show as PostsShow } from './pages/posts/Show';
import { Edit as PostsEdit } from './pages/posts/Edit';

import './App.css';

function App() {

  const helmetContext = {};

  return (
    <BrowserRouter>
      <HelmetProvider context={helmetContext}>
      <Header />
      <div className="content">
        <Routes>
          <Route path="/" element={<Home/>} />
          <Route path="/posts" element={<PostsList/>}/>
          <Route path="/posts/new" element={<PostsNew/>} />
          <Route path="/posts/:id" element={<PostsShow/>} />
          <Route path="/posts/:id/edit" element={<PostsEdit/>} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </div>
      </HelmetProvider>
    </BrowserRouter>
  );
}

export default App;
