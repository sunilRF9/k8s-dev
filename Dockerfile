FROM ubuntu

COPY ./k8sdev ./k8sdev

CMD [ "./k8sdev" ]
