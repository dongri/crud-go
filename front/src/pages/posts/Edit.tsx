import { useState, useEffect } from 'react'
import { Link, useParams, useNavigate } from 'react-router-dom'
import { http } from '../../api/http';
import { Loading } from '../../components/Loading';
import { Head } from '../../components/Head'

interface PostsResponse {
  post: Post
}

interface Post {
  id:      number,
	title:   string,
  body:    string,
  created: string,
  updated: string,
}

export const Edit = () => {
  const params = useParams();
  const navigate = useNavigate();

  const [error, setError] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);

  const [post, setPost] = useState<Post>({} as Post);

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

  useEffect(() => {
    const getPosts = async () => {
      try {
        setLoading(true);
        const response = await http<PostsResponse>(`/posts/${params.id}`, "GET");
        const body = response.parsedBody
        if (!body) {
          setError('response error')
          return
        }
        setPost(body.post);
        setState(body.post);
      } catch (err: any) {
        setError(err.toString());
      } finally {
        setLoading(false);
      }
    }
    getPosts();
  },[params]);

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await http<PostsResponse>(`/posts/${params.id}/update`, "POST", {
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

  const onDelete = async () => {
    try {
      const response = await http<PostsResponse>(`/posts/${params.id}/delete`, "POST", {});
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
      <Head title="Edit" description="edit"/>
      <div>
        <h2>Posts</h2>
        <div>{error}</div>
        <form onSubmit={onSubmit}>
          <table>
            <tbody>
              <tr>
                <th>Title</th>
                <td><input type="text" name="title" defaultValue={post.title} onChange={handleChange} /></td>
              </tr>
              <tr>
                <th>Body</th>
                <td><textarea name="body" defaultValue={post.body} onChange={handleChange} /></td>
              </tr>
            </tbody>
          </table>
          <div className='v1'></div>
          <Link to={`/posts/${params.id}`}>Cancel</Link>
          <span className='h1'>|</span>
          <button type="submit">Update</button>
          <span className='h1'>|</span>
          <button onClick={onDelete}>Delete</button>
        </form>
        <Loading loading={loading} />
      </div>
    </>
  )
}
