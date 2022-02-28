FROM centos

RUN cd /etc/yum.repos.d/ \
    && sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-* \
    && sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

RUN yum install golang -y 
RUN yum install dlv -y 
RUN yum install binutils -y 
RUN yum install vim -y 
RUN yum install gdb -y 

# docker build -t myos .
# docker run -it --rm myos bash