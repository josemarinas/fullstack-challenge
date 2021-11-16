import { ImageUploaderProps } from "./ImageUploader.props";
import { create } from 'ipfs-http-client'
import { Button, Card, CardContent, Typography, Grid } from '@mui/material'
import { useState } from "react";
import { useMutation } from 'react-query'
import axios from "axios";
import { ApiImage } from "../../types"
import { Conf } from "../../conf";

export function ImageUploader(props: ImageUploaderProps) {
  const [lastUUID, setLastUUID] = useState<Array<string>>([])
  const [error, setError] = useState<any>()
  const createImage = useMutation((newImage: ApiImage) => {
    return axios.post(Conf.baseUrl + '/image', newImage)
  })
  const ipfs = create({ url: Conf.ipfsUrl })
  const uploadImages = async (files: FileList) => {
    for (let i = 0; i < files.length; i++) {
      // select file
      const file = files[i]
      // setup details
      const fileDetails = {
        path: file.name,
        content: file
      }
      const options = {
        wrapWithDirectory: false,
      }
      try {
        const { cid } = await ipfs.add(fileDetails, options)
        await ipfs.pin.add(cid)
        const { data } = await createImage.mutateAsync({ cid: cid.toString() })
        setLastUUID(data.id)
        setError(null)
      } catch (e) {
        setError(e)
      }
    }
  }
  return (
    <Card>
      <CardContent>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <Typography sx={{ fontSize: 32, fontWeight: 800 }}>
              Image Uploader
            </Typography>
          </Grid>

          <Grid item xs={12}>
            <Typography variant="body1" component="div">
              Last Upload: {lastUUID}
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <Button
              variant="contained"
              component="label"
            >
              Upload File
              <input
                type="file"
                accept="image/*"
                hidden
                multiple={true}
                onChange={(e) => e.target.files ? uploadImages(e.target.files) : ''}
              />
            </Button>
          </Grid>
          <Grid item xs={12}>
            {error ? 'Error:' + error.message : ''}
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  )
}