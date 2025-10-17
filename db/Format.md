## .mercury File format

`.mercury` files are lists of blocks in the following format

| Type        | Size     | Description                                                        |
| ----------- | -------- | ------------------------------------------------------------------ |
| string      | 7 bytes  | A magic number spelling out `MERCURY`                            |
| UUID        | tbd      | A UUID version tbd thus size is TBD                                |
| int32       | 4 bytes  | The size of the data block to be read                              |
| gzipped gob | Variable | The data block that has been compressed with gzip level `<tbd>` |

The first data block describes the geometry of the data and after that is meta data for the geometry
