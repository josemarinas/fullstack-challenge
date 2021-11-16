import { ImageViewerProps } from "./ImageViewer.props";
import axios from "axios";
import { create } from 'ipfs-http-client'
import { ChangeEvent, useState } from "react";
import { Conf } from "../../conf";
import { Button, Card, CardContent, Typography, CardMedia, TextField, Grid } from '@mui/material'
import { IsValidUUIDV4 } from "../../utils/validators";
import { uint8chunksToBase64 } from "../../utils/bytes";

export function ImageViewer(props: ImageViewerProps) {
  const [imageUUID, setImageUUID] = useState<string>('')
  const [imageCid, setImageCid] = useState<string | null>(null)
  const [imageSrc, setImageSrc] = useState<string>('')
  const [error, setError] = useState<any>()
  const [textFieldHasError, setTextFieldHasError] = useState<any>(false)
  const ipfs = create({ url: Conf.ipfsUrl })
  const getImage = async () => {
    try {
      const { data } = await axios.get(Conf.baseUrl + '/image/' + imageUUID)
      setImageCid(data.cid)
      const chunks = []
      for await (const chunk of ipfs.cat(data.cid)) {
        chunks.push(chunk)
      }
      setImageSrc(uint8chunksToBase64(chunks))
      setError(null)
    } catch (e) {
      setError(e)
    }
  }
  const handleTextFieldChange = (e: ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
    const value = e.target.value
    setTextFieldHasError(!IsValidUUIDV4(value))
    setImageUUID(value)
  }
  return (
    <Card>
      <CardContent>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <Typography sx={{ fontSize: 32, fontWeight: 800 }}>
              Image Viewer
            </Typography>
          </Grid>
          {imageCid ?
            <Grid item xs={12}>
              <Typography variant="body1" component="div">
                <p>Image CID: {imageCid}</p>
                <a rel="noreferrer" target="_blank" href={'https://ipfs.io/ipfs/' + imageCid}>View on IPFS gateway</a>
              </Typography>
            </Grid> : ''}
          <Grid item xs={12}>
            <TextField
              fullWidth
              error={textFieldHasError}
              helperText={textFieldHasError ? 'Invalid UUID' : ''}
              label="UUID"
              variant="outlined"
              onChange={handleTextFieldChange}
            />
          </Grid>
          <Grid item xs={12}>
            <Button
              disabled={textFieldHasError}
              variant="contained"
              component="label"
              onClick={getImage}
            >
              Get File
            </Button>
          </Grid>
          <Grid item xs={12}>
            {error ? 'Error:' + error.message : ''}
          </Grid>
          {
            imageCid ?

              <Grid item xs={12}>
                <CardMedia height="500" width="500" component='img' src={'data:image/png;base64,' + imageSrc} />
              </Grid>
              : ''
          }
        </Grid>
      </CardContent>
    </Card>
  )
}
