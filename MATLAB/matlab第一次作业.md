# 第一次作业



## 程序

- mToNCost.m(m到n-1期的成本)

```matlab
function cost = mToNCost(d,k,c,h,m,n)
%mToNCost - 计算第m期生产满足第m到n-1期的所有需求带来的成本
%
%输入：
%   d(vector) : 各阶段的需求
%   k(vector) : 各阶段的固定成本
%   c(vector) : 各阶段的单位边际成本
%   h(vector) : 各阶段的持有库存的边际成本
%   m(number) : 开始的阶段
%   n(number) : 结束的后一阶段
%
%输出：
%   cost(number) : 第m期生产满足第m到n-1期的所有需求带来的成本

cost = k(m) + c(m) .* sum(d(m:n-1));
% 遍历从m 到 n-1期的 所有 holding cost 
for i = m:n-2
    cost = cost + h(i) .* sum(d(i+1:n-1));
end

% 第二种实现方式
% \[ c_{m,i} \cdot d_i = c_m * d_i + (h_m + \cdots + h_{i-1}) * d_i  \]
cost = k(m);
for i = m:n-1
   cost = cost + c(m) .* d(i) + sum(h(m:i-1)) .* d(i);
end
end
```

- 根据s(n)迭代求出最优路线

```matlab
function road = dyRoad(leng,s)
%dyRoad - 根据动态规划给出的s，递归的给出最优决策方案
%
%路径的递归定义： 1 ---> ... ---> s(s(n) - 1) --->  s(n) ---> N 
%输入：
%   leng(number):阶段数目
%   s(vector):在动态规划算法中给出的s
%输出：
%   road(vector) : 达到最小成本的方案(0代表不生产，1代表生产)

road = zeros(1,leng);
point = leng;
while point > 0 
    road(s(point)) = 1;
    point = s(point) - 1;
end
end
```

- 主程序

```matlab
function [result,road] = dySolution(d,k,c,h)
%dySolution - 利用自底向上的动态规划算法求解DELS问题
%
%算法时间复杂度为O(n^2),n为维度
%输入：
%   d(vector) : 各阶段的需求
%   k(vector) : 各阶段的固定成本
%   c(vector) : 各阶段的单位边际成本
%   h(vector) : 各阶段的持有库存的边际成本
%
%输出：
%   result(number) : 最小成本
%   road(vector) : 达到最小成本的方案(0代表不生产，1代表生产)
%
%example(d,k,c,h都为n维向量):
%   [optResult,road] = dySolution(d,k,c,h)

% 验证输入向量的维度
if any([size(d)~=size(k),size(d)~=size(c),size(d)~=size(h)]) 
    disp('请输入维度相同的向量');
    result = -1;
    road = -1;
    return 
end

% 预分配内存
% r(n) 表示从第1期到第n期的最小成本，该问题是动态规划的子问题。
% s(n) 表示r(n)对应最优子方案中最后一次生产的时期是在 第s(n)阶段。
r = zeros(1,length(d));
s = zeros(1,length(d));

% 遍历求解子问题
% 所有的子问题 都只计算了一次，算法的时间复杂度为 O(n^2) 
for i = 1:length(d)
    r(i) = mToNCost(d,k,c,h,1,i+1);
    s(i) = 1;
    for j = 1:i-1
        temp = r(j) + mToNCost(d,k,c,h,j+1,i+1);
        if temp < r(i) 
            r(i) = temp;
            s(i) = j+1;
        end
    end
end
% 最小成本
result = r(length(d));
% 根据s(n)计算路径 
road = dyRoad(length(d),s);
end
```



- 检查是否是最优的路线

```matlab
function result = checkOptRoad(d,k,c,h,road)
%checkOptRoad - 判断给定路径是否是动态规划的最优解
%
%输入：
%   d(vector) : 各阶段的需求
%   k(vector) : 各阶段的固定成本
%   c(vector) : 各阶段的单位边际成本
%   h(vector) : 各阶段的持有库存的边际成本
%
%输出：
%   result(boolean) : 逻辑1或者逻辑0

plan = find(road == 1);
sum = 0;
for i = 1:length(plan)-1
    sum = sum + mToNCost(d,k,c,h,plan(i),plan(i+1));
end
sum = sum + mToNCost(d,k,c,h,plan(length(plan)),length(d)+1);
% disp(sum);
% disp(dySolution(d,k,c,h));
result = sum == dySolution(d,k,c,h);
end
```



- 单元测试，压力测试，回到龙爸的问题

```matlab
% 压力测试
stressD = randi(100,1000,10);
stressK = randi(100,1000,10);
stressC = randi(100,1000,10);
stressH = randi(100,1000,10);
tic
for i = 1:1000
    dySolution(stressD(i,:),stressK(i,:),stressC(i,:),stressH(i,:));
end
fprintf('压力测试平均耗时:%d\n',toc/1000);

% 测试第一问中的结论
d = randi(10,10,10);
k = randi(10,10,10);
c = randi(10,10,10);
h = randi(10,10,10);
tem = rand();
for i = 1:10
    [a1,b1] = dySolution(d(i,:),k(i,:),c(i,:),h(i,:));
    [a2,b2] = dySolution(d(i,:),k(i,:) + tem,c(i,:),h(i,:));
    [a3,b3] = dySolution(d(i,:),k(i,:),c(i,:) + tem,h(i,:));
    [a4,b4] = dySolution(d(i,:),k(i,:)+tem,c(i,:)+tem,h(i,:));
    if ~all([checkOptRoad(d(i,:),k(i,:),c(i,:) + tem,h(i,:),b1),checkOptRoad(d(i,:),k(i,:) + tem,c(i,:),h(i,:),b1),checkOptRoad(d(i,:),k(i,:)+tem,c(i,:)+tem,h(i,:),b1)]) 
        disp('测试不通过');
    end
end

% 任务2中第一问结论正确
d = randi(10,100,10);
k = zeros(1,10);
c = randi(10,1,10);
h = randi(10,1,10);

result = zeros(100,10);
for i = 1:100
    [a1,b1] = dySolution(d(i,:),k,c,h);
    result(i,:) = b1;
end

% 任务三的更新过程已经蕴含在 动态规划s(n)中了


```

