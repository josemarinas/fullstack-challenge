import Pages from "../types"
import { ImageViewer, ImageUploader } from '../components'
import { Grid } from "@mui/material"
type Props = {
  onSetPage: (page: Pages) => void
}
export function HomePage(props: Props) {
  return (
    <Grid container spacing={2}>
      <Grid item xs={6}>
        <ImageUploader></ImageUploader>
      </Grid>
      <Grid item xs={6}>
        <ImageViewer></ImageViewer>
      </Grid>
    </Grid>
  )
}