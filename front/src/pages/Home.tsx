import { Link } from 'react-router-dom'
import { Head } from '../components/Head'

export const Home = () => {

  return (
    <>
      <Head title="Home" description="home" />
      <div>
        <h2>Home</h2>
        <Link to={`/posts/`}>Posts</Link>
      </div>
    </>
  )
}
