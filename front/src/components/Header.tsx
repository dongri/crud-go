import { Link } from "react-router-dom"

export const Header = () => {

  return (
    <div className="header">
      <div className="header-left">
          <h1>
            <Link to={`/`} >crud-go</Link>
          </h1>
        </div>
      <div className="header-right">
      </div>
    </div>
  )
}
