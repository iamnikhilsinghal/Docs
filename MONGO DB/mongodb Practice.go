use nikdb
db.createCollection('nikcoll')
show collections
db.nikcoll.insert({first_name:"nikhil", last_name:"singhal"})
db.nikcoll.find()
db.nikcoll.insert([{first_name:"nikhil", last_name:"singhal"},{first_name:"amit",last_name:"upadhyay"}])		//Multiple insertion
db.nikcoll.find()
db.nikcoll.find().pretty()
db.nikcoll.insert({first_name:"suraj", last_name:"pathak",gender:"male"})
db.nikcoll.update({first_name:"suraj"},{first_name:"suraj",last_name:"pathak",sex:"male"})						//Long Process
db.nikcoll.update({first_name:"amit"},{$set:{gender:"male"}})													// Using $set
db.nikcoll.update({first_name:"amit"},{$set:{age:45}})
db.nikcoll.update({first_name:"amit"},{$inc:{age:5}})															//Increment
db.nikcoll.update({first_name:"amit"},{$unset:{age:1}})															//Remoce column
db.nikcoll.update({first_name:"sudarshan"},{first_name:"sudarshan",last_name:"prajapati"})						//Will show writeresult but don't insert
db.nikcoll.update({first_name:"sudarshan"},{first_name:"sudarshan",last_name:"prajapati"},{upsert:true})		//Will insert now
db.nikcoll.update({first_name:"amit"},{$rename:{"gender":"sex"}})												//Remove column
db.nikcoll.remove({first_name:"nikhil"})																		//All doc with name "nikhil"	
db.nikcoll.remove({first_name:"nikhil"},{justone:true})															//Will work with objectID only 
db.nikcoll.find({first_name:"nikhil"})																			//All "nikhil"
db.nikcoll.find({$or:[{first_name:"nikhil"},{first_name:"suraj"}]})												//All "nikhil" OR "suraj"
db.nikcoll.find({age:{$lt:40}}).pretty()
db.nikcoll.find().sort({last_name:1})
db.nikcoll.find().sort({last_name:-1})
db.nikcoll.find().count()
db.nikcoll.find({gender:"male"}).count()
db.nikcoll.find().limit(4)																						//First 4
db.nikcoll.find().limit(4).sort({last_name:1})
db.nikcoll.find().forEach(function(doc){print("Customer name:"+doc.first_name)})							


db.employee.aggregate([{$group : {_id : "$firstName", countOfEmp : {$sum : 1}}}])	select firstName, count(*) countOfEmp from employee group by firstName;
db.employee.aggregate([{$match:{salary:{$gt:2000}}},{$count:"firstname"}])			select count(*) from employee where salary >= 2000
db.employee.aggregate([{$group:{_id:"$firstName", salary_sum:{$sum:"$salary"}}}])	select firstName, sum(salary) salary_sum from employee group by firstName;
db.employee.aggregate([{$group:{_id:null, salary_sum:{$sum:"$salary"}}}])			select sum(salary) salary_sum from employee;
db.employee.aggregate([{ $match: { salary : { $gt: 2000} } },						select firstName, sum(salary) salary_sum from employee 
 			{$group : {_id : "$firstName", salary_sum : {$sum : "$salary"}}}])				where salary > 2000 group by firstName
db.employee.aggregate([{$group: { _id: "$firstName", salary_first:{$first:"$salary"}}}])	first salary of employee with same firstName	
db.employee.aggregate([{$group: {_id : "$firstName", salary_last : {$last : "$salary"}}}])	Last salary of employee with same firstName
db.employee.aggregate([{$group: { _id: "$firstName",all_salary:{$push:"$salary"}}}])		all salary of all employee with same firstName in array


https://www.javamadesoeasy.com/2017/03/last-operator-in-mongodb.html