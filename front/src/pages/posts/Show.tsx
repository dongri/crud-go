import { useState, useEffect } from 'react'
import { Link, useParams } from 'react-router-dom'
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

export const Show = () => {
  const params = useParams();

  const [error, setError] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);

  const [post, setPost] = useState<Post>({} as Post);

  useEffect(() => {
    const getPosts = async () => {
      try {
        setLoading(true);
        const id = params.id;
        const response = await http<PostsResponse>(`/posts/${id}`, "GET");
        const body = response.parsedBody
        if (!body) {
          setError('response error')
          return
        }
        setPost(body.post);
      } catch (err: any) {
        setError(err.toString());
      } finally {
        setLoading(false);
      }
    }
    getPosts();
  },[params]);

  return (
    <>
      <Head title="Show"  description="show"/>
      <div>
        <h2>Posts</h2>
        <div>{error}</div>
        <table>
            <tbody>
              <tr>
                <th>Title</th>
                <td>{post.title}</td>
              </tr>
              <tr>
                <th>Body</th>
                <td>{post.body}</td>
              </tr>
              <tr>
                <th>Created</th>
                <td>{post.created}</td>
              </tr>
              <tr>
                <th>Updated</th>
                <td>{post.updated}</td>
              </tr>
            </tbody>
          </table>
        <div className='v1'></div>
        <Link to={`/posts`}>Back</Link>
        <span className='h1'>|</span>
        <Link to={`/posts/${post.id}/edit`}>Edit</Link>
        <Loading loading={loading} />
      </div>
    </>
  )
}
