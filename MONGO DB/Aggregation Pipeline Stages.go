db.createCollection("scores")
db.scores.insert([{_id:1,student:"Maya",homework:[10,5,8]},{_id:2,student:"Ramu",homework:[9,41,5]}])

$addFields:
db.scores.aggregate([{$addFields: {totalhw:{$sum:"$homework"}}}])	//Will add--> "totalhw" : 23 
db.scores.aggregate([{$addFields: {"student":"Golang"}}])			//will change student names with Golang

$count:
db.scores.insert([{ _id : 1, score : 88 },{ _id : 2, score : 78 },{ _id : 3, score : 97 },{ _id : 4, score : 71 }])
db.scores.aggregate([{$match:{"score":{$gt:80}}},{$count:"Scores>80"}])		//"Scores>80" : 2

$graphLookup: operation recursively matches on the reportsTo and name fields in the employees collection, 
			  returning the reporting hierarchy for each person.

$group:
db.scores.insert([{ _id : 1, score : 88 },{ _id : 2, score : 78 },{ _id : 3, score : 97 },{ _id : 4, score : 71 }])
db.scores.aggregate([{$group: { _id: "$subject", totalscore:{$sum:"$score"}}}])		//{	"_id" : "History",	"totalscore" : 166} .........

$limit:
db.scores.aggregate({$limit: 2})

$match: Simplest query:
db.scores.aggregate({$match: {"subject":"hindi"}})

$project:
db.scores.aggregate({$project: {_id:0, score:1}})	//only score 
db.scores.aggregate({$project: {score:1},{$match:score:{$gt:80}}}) Problem so use "$eq

$set:
db.scores.aggregate({$set: { "class":10 }})		//Add class:10 in every record

$unset:
db.scores.aggregate([ { $unset: "homework" } ])		//Remove "homework" field

$skip
db.article.aggregate({ $skip : 5 })			//skips the first 5 documents

$sort:
db.users.aggregate([{ $sort : { age : -1, posts: 1 } }]) //descending order of age and ascending order of posts




