FROM crystallang/crystal:0.30.1

COPY . /src
WORKDIR /src
RUN crystal build -o /out --link-flags "-static" main.cr

CMD cat /out
