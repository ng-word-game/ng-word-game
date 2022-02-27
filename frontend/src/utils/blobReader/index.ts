const blobReader = (blob: Blob) => {
  const fileReader = new FileReader()

  return new Promise((resolve, reject) => {
    fileReader.onerror = () => {
      fileReader.abort()
      // eslint-disable-next-line prefer-promise-reject-errors
      reject()
    }

    fileReader.onload = () => {
      resolve(fileReader.result as string)
    }

    fileReader.readAsText(blob)
  })
}

export const blobToJson = async (blob: Blob) => {
  const blobtext = await blobReader(blob) as string
  return JSON.parse(blobtext)
}
