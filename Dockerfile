FROM scratch
ARG OS
ARG ARCH
ADD build/${OS}_${ARCH}/demoapp /demoapp
CMD ["/demoapp"]
