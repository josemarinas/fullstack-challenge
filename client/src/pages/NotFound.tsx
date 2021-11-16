import Pages from "../types"
type Props = {
  onSetPage: (page:Pages) => void 
}
export function NotFoundPage (props:Props) {
  return (
    <div>
      <h1>
        Whops! Not Found
      </h1>
      <button onClick={() => props.onSetPage(Pages.Home)}>
        Pls go back
      </button>
    </div>
  )
}