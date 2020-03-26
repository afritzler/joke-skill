# joke-skill

Random Joke generator skill for SAP CAI

## Deploy Cloud Function

```shell script
gcloud functions deploy joke-skill --runtime go113 --trigger-http --region europe-west3 --memory 128MB --entry-point RandomJoke
```
