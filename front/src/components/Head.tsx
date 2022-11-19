import {FC} from "react";
import { Helmet } from "react-helmet-async"

type HeadProps = {
  title: string | "crud-go",
  description: string | "crud-go",
}

export const Head: FC<HeadProps> = (props) => {
  const { title, description } = props

  const domain = 'https://front-vjiy.onrender.com'

  return (
    <Helmet
      title={ title }
      meta={[
        { name: 'description', content: description },
        { property: 'og:title', content: title },
        { property: 'og:type', content: 'website' },
        { property: 'og:url', content: domain + '/about' },
        { property: 'og:image', content: domain + '/favicon.ico' },
        { property: 'og:description', content: description },
        { name: 'twitter:card', content: 'summary_large_image' },
        { name: 'twitter:image', content: domain + '/favicon.ico' },
        { name: 'twitter:site', content: title },
      ]}
    />
  )
}
