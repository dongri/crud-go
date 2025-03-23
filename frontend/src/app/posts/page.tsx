'use client';

import { useEffect, useState } from 'react';
import { PublicGet } from '@/api/client';
import { Post } from '@/types/post';
import Link from 'next/link'

export default function Index() {

  const [posts, setPosts] = useState<Post[]>([]);

  const [isLoading, setIsLoading] = useState(true)

  const fetchPosts = async () => {
    setIsLoading(true);
    const result = await PublicGet('/posts');
    if (result.error) {
      alert(result.error);
      return
    }
    setPosts(result.posts);
    setIsLoading(false);
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  return (
    <div>
      <h1 className='text-2xl font-bold'>Posts</h1>
      <div className='mt-4 mb-4'>
        {isLoading ? (
          <p>Loading posts...</p>
        ) : posts.length === 0 ? (
          <div>
            <p>No posts found</p>
          </div>
        ) : (
          <ul className='list-disc list-inside'>
            {posts.map((post) => (
              <li key={post.id}>
                <Link href={`/posts/${post.id}`}>{post.title}</Link>
              </li>
            ))}
          </ul>
        )}
      </div>
      <div>
        <Link href="/posts/new" className='text-blue-500'>New Post</Link>
      </div>
    </div>
  );
}
