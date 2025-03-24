'use client';

import { useEffect, useState } from 'react';
import { useClient } from '@/api/client';
import { Post } from '@/types/post';
import Link from 'next/link'

export default function Index() {

  const [posts, setPosts] = useState<Post[]>([]);

  const { publicGet } = useClient();

  const fetchPosts = async () => {
    const result = await publicGet('/posts') as { posts: Post[] }
    setPosts(result.posts);
  };

  useEffect(() => {
    fetchPosts();
  }, []); // eslint-disable-line

  return (
    <div>
      <h1 className='text-2xl font-bold'>Posts</h1>
      <div className='mt-4 mb-4'>
        {posts.length === 0 ? (
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
