# 毛康力

## 基本信息

<table>
	<tr>
		<td>毕业学校：华中科技大学</td>
		<td>专业：计算机软件与理论</td>	
	</tr>
	<tr>
		<td>学历：硕士</td> 	
		<td>性别：男</td>
	</tr>
	<tr>
		<td>手机：18520140924</td>
		<td>邮箱：tiancaiamao@gmail.com</td>
	</tr>
</table>

技术博客：[http://www.zenlife.tk](http://www.zenlife.tk)

## 工作经验

2014.7~至今 广州舜飞信息科技有限公司

基础建设层面：

1. 性能优化

广告竞价服务是一个CPU密集型的业务，又是一个延迟要求很苛刻的场景。在高并发压力下面，CPU过高一方面会消耗硬件资源，另一方面会影响垃圾回收，继而导致响应超时而失败。刚入职的半个月，为系统做过一个性能优化，使竞价系统的CPU降为原来53%，响应时间也大大缩短。

2. kv存储

cookie mapping的存储方案要求高并发，低延迟。之前是用redis实现，业务层做切分。存在的问题是数据量大，成本高，可维护性差。在调研SSD的可行性之后，推荐公司采用了另一套开源系统的方案，使成本降低至少6倍以上，并且可维护性提高。其中一个集群目前存储数T的数据，100亿+记录，稳定运行1年以上。 

3. 配置中心

大多项目配置使用json本地文件，分散在各个机器上面。集群的节点到一定规模以后，配置的维护和管理很不方便。于是基于consul实现了配置中心。提供了友好的UI，保证配置的一致性，高可用性。支持版本回滚，更新会近实时地通知到注册业务。

4. 文件推送系统

有些业务有加载数据的需求，比如算法团队的一些数据要加载到竞价进程中提供决策，竞价进程分散在很多机器。这个系统的功能是将指定的文件推送到一批机器。做了P2P的数据传输实现，并提供进度以及成功失败相关的监控信息，外围还有一些更新回调的通知机制。

5. 业务监控/日志收集

做过一些相关的方面调研工作，但没有用上。

业务层面：

* DSP广告竞价系统

每日100亿+PV的系统，也是公司最核(ying)心(li)的业务。除了性能优化，还负责功能改进，渠道对接，以及产品需求方面的开发。

* 动态创意系统

主导了整体的架构设计，大部分框架搭建和开发，协调其它成员的工作，以及部门间的沟通。

* TagManager项目

CURD类型的小项目，独立完成开发和负责维护工作。

## 个人项目

* cora 一个编译器，从scheme语言生成x86的汇编。

喜欢scheme语言，并且对汇编/编译原理/运行时都有比较深入的研究，所以写这个东西。

写一个编译器涉及到的知识量巨大！从前期的CPS变换，闭包转换；到后期的寄存器分配，编译优化；以及运行时的设计，垃圾回收等，非常的锻炼人。

* 《深入解析Go》一本开源的电子书。

对Go语言底层实现的东西比较感兴趣，并且热爱分享知识，所以写了这本书。

书中涵盖了Go语言底层实现的很多方面，从基本的数据类型实现，到汇编代码，内存管理，垃圾回收，调度等等。还有一些内容有待完善，继续更新中。

目前github的star数500+，fork数128。

项目地址：[https://github.com/tiancaiamao/go-internals](https://github.com/tiancaiamao/go-internals)