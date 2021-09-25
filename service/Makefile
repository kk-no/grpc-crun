.PHONY: deploy
deploy:
	gcloud run deploy service \
		--source . \
		--region asia-northeast1 \
		--platform managed \
		--service-account cloud-run@sandbox-kkeino.iam.gserviceaccount.com
