# Media module

## Module purpose

The media module is mainly responsible for unified management of all media resources in the community.

## Module scope

This module is used to uniformly manage resources in cloud storage in projects.

- Input: resources
- Output: ID after adding the resource record
- Dependencies: Cloud Storage

## Module structure
```go
type File struct {
	Id       int
	FileKey  string
	Format   string
	UserId   int
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

```


## Module Interface

None

## Functions

### Upload resources

- Function: Upload resources
- Input: resources
- Return: media table id
- Error: Cloud storage and database

func SaveMedia
```go
func (c *Community) SaveMedia(ctx context.Context, userId string, data []byte) (int64, error)
```

### Access to resources

- Function: Get the URL of the resource
- Input: media id
- Return: media table fileKey
- Error: The resource may not exist, returning 404

func getMediaInfo
```go
func (c *Community) getMediaInfo(fileKey string) (*File, error)
```