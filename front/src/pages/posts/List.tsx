import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { http } from '../../api/http';
import { Loading } from '../../components/Loading';
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

export const List = () => {
  const [error, setError] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);

  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    const getPosts = async () => {
      try {
        setLoading(true);
        const response = await http<PostsResponse>(`/posts`, "GET");
        const body = response.parsedBody
        if (!body) {
          setError('response error')
          return
        }
        setPosts(body.posts);
      } catch (err: any) {
        setError(err.toString());
      } finally {
        setLoading(false);
      }
    }
    getPosts();
  },[]);

  return (
    <>
      <Head title="List" description="list" />
      <div>
        <h2>Posts</h2>
        <div>{error}</div>
        <Link to={`/posts/new`}>New Post</Link>
        <div>
          {posts.map(post => (
            <div key={post.id} className="post">
              <h4>{post.title}</h4>
              <div>
                {post.body}
              </div>
              <div className="v1">
                {post.created}
              </div>
              <div className="v1"></div>
              <Link to={`/posts/${post.id}`}>Show</Link>
            </div>
          ))}
        </div>
        <Loading loading={loading} />
      </div>
    </>
  )
}
