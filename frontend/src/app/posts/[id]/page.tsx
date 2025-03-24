'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';
import { useRouter, useParams } from 'next/navigation';

import { useClient } from '@/api/client';
import { Post } from '@/types/post';
import { Comment } from '@/types/comment';

export default function Show() {
  const router = useRouter();
  const { publicGet, publicPost } = useClient();

  const params = useParams();

  const id = params?.id;

  const [post, setPost] = useState<Post>();

  const [comments, setComments] = useState<Comment[]>([]);

  const [state, setState] = useState({
    body: "",
  });

  const fetchPost = async () => {
    const result = await publicGet(`/posts/${id}`);
    if (result.error) {
      alert(result.error);
      return
    }
    setPost(result.post);
  };

  const fetchComments = async () => {
    const result = await publicGet(`/posts/${id}/comments`);
    if (result.error) {
      alert(result.error);
      return
    }
    setComments(result.comments);
  };

  useEffect(() => {
    fetchPost();
    fetchComments();
  }, []); // eslint-disable-line

  const handleInputChange = async (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    if (event.target.type === "number") {
      setState({ ...state, [event.target.name]: Number(event.target.value) });
    } else {
      setState({ ...state, [event.target.name]: event.target.value });
    }
  };

  const handleSubmit = async () => {
    try {
      const data = await publicPost(`/posts/${id}/comments`, {
        body: state.body,
      });

      if (data.error) {
        alert(data.error);
        return
      }
      state.body = "";
      fetchComments();
    } catch (error) {
      console.error('Error submitting sponsorship:', error);
      alert('An error occurred while submitting the sponsorship. Please try again.');
    } finally {
    }
  }

  const handleDelete = async () => {
    try {
      const data = await publicPost(`/posts/${id}/delete`, {});

      if (data.error) {
        alert(data.error);
        return
      }
      alert('Post deleted');
      router.push('/posts');
    }
    catch (error) {
      console.error('Error deleting post:', error);
      alert('An error occurred while deleting the post. Please try again.');
    }
  }

  return (
    <div>
      <h1 className='text-2xl font-bold'>Post</h1>
      <div className='mt-4 mb-4'>
        <div>
          {post?.title}
        </div>
        <div>
          {post?.body}
        </div>
      </div>
      <div>
        <Link href={`/posts/${id}/edit`} className='text-blue-500 m-2'>Edit</Link>
        <Link href="/posts" className='text-blue-500 m-2'>Back</Link>
        <Link href="#" className='text-blue-500 m-2' onClick={handleDelete}>Delete</Link>
      </div>
      <div className='mt-4 mb-4'>
        <h2 className='text-l font-bold'>Comments</h2>
        <ul className='list-disc list-inside'>
          {comments.map((comment) => (
            <li key={comment.id}>
              {comment.body}
            </li>
          ))}
        </ul>

        <form>
          <textarea
            className='border border-gray-300 rounded p-2'
            name="body"
            value={state.body}
            onChange={handleInputChange}
            required
          />
          <div>
            <button
              className='bg-blue-500 text-white rounded p-2 mt-4'
              type="button"
              onClick={handleSubmit}
            >
              Submit
            </button>
          </div>
        </form>

      </div>
    </div>
  );
}
