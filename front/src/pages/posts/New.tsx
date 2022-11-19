import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { http } from '../../api/http';
import { Head } from '../../components/Head'

interface PostsResponse {
  posts: Post[]
}

interface Post {
  id:      number,
	title:   string,
  body:    string,
  created: string,
  updated: string,
}

export const New = () => {
  const navigate = useNavigate();

  const [error, setError] = useState<string>("");

  const [state, setState] = useState<Post>({
    id: 0,
    title: "",
    body: "",
    created: "",
    updated: "",
  });

  const handleChange = async (event: any) => {
    if (event.target.type === "number") {
      setState({ ...state, [event.target.name]: Number(event.target.value) });
    } else {
      setState({ ...state, [event.target.name]: event.target.value });
    }
  };

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await http<PostsResponse>(`/posts/new`, "POST", {
        title:   state.title,
        body:    state.body,
      });
      const resBody = response.parsedBody
      if (!resBody) {
        setError('response error')
        return
      }
      navigate('/posts');
    } catch (err: any) {
      setError(err.toString());
    }
  }

  return (
    <>
      <Head title="New" description="new" />
      <div>
        <h2>Posts</h2>
        <div>{error}</div>
        <form onSubmit={onSubmit}>
          <table>
            <tbody>
              <tr>
                <th>Title</th>
                <td><input type="text" name="title" defaultValue={""} onChange={handleChange} /></td>
              </tr>
              <tr>
                <th>Body</th>
                <td><textarea name="body" onChange={handleChange} /></td>
              </tr>
            </tbody>
          </table>
          <div className='v1'></div>
          <Link to={`/posts`}>Cancel</Link>
          <span className='h1'>|</span> 
          <button type="submit">Post</button>
        </form>
      </div>
    </>
  )
}
