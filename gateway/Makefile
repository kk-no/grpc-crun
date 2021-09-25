.PHONY: deploy
deploy:
	gcloud run deploy gateway \
		--source . \
		--region asia-northeast1 \
		--platform managed \
		--service-account cloud-run@sandbox-kkeino.iam.gserviceaccount.com \
		--set-env-vars "SERVICE_DOMAIN=service-vj7vnahacq-an.a.run.app" \
		--set-env-vars "SERVICE_PORT=443"
