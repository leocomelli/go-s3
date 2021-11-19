# go-s3

A simple implementation to upload file to AWS S3.

### Environment variables:

* AWS_S3_REGION
* AWS_S3_OBJECT_PATH
* AWS_S3_BUCKET
* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY

### Usage:

```bash
export AWS_S3_REGION=us-east-1
export AWS_S3_OBJECT_PATH=myfile.txt
export AWS_S3_BUCKET=no-public-bucket
export AWS_ACCESS_KEY_ID=my-access-key
export AWS_SECRET_ACCESS_KEY=my-secret-key

./go-s3
```

### Output:

```bash
Location: https://no-public-bucket.s3.amazonaws.com/myfile.txt
*** no-public-bucket ***
-  myfile.txt

```

