'use client';

import { useState } from 'react';
import { useClient } from '@/api/client';
import { useRouter } from 'next/navigation'
import Link from 'next/link';

export default function New() {
  const router = useRouter();
  const { publicPost } = useClient();

  const [state, setState] = useState({
    title: "",
    body: "",
  });

  const handleInputChange = async (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    if (event.target.type === "number") {
      setState({ ...state, [event.target.name]: Number(event.target.value) });
    } else {
      setState({ ...state, [event.target.name]: event.target.value });
    }
  };

  const handleSubmit = async () => {
    try {
      const data = await publicPost('/posts', {
        title: state.title,
        body: state.body,
      });

      if (data.error) {
        alert(data.error);
        return
      }
      router.push('/posts');
    } catch (error) {
      console.error('Error submitting sponsorship:', error);
      alert('An error occurred while submitting the sponsorship. Please try again.');
    } finally {
    }
  }

  return (
    <div>
      <h1 className='text-2xl font-bold'>New Post</h1>
      <form>
        <div>
          <label className='block'>Title</label>
          <input
            className='border border-gray-300 rounded p-2'
            id="title"
            name="title"
            type="text"
            value={state.title}
            onChange={handleInputChange}
            required
          />
        </div>
        <div>
          <label className='block'>Body</label>
          <textarea
            className='border border-gray-300 rounded p-2'
            name="body"
            value={state.body}
            onChange={handleInputChange}
            required
          />
        </div>
        <div>
          <button
            className='bg-blue-500 text-white rounded p-2 mt-4'
            type="button"
            onClick={handleSubmit}
          >
            Submit
          </button>
          <Link href="/posts" className='text-blue-500 m-2'>Cancel</Link>
        </div>
      </form>

    </div>
  );
}
