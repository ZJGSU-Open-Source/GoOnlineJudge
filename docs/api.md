#API

* log need to be recorded automatically by server.
* all the users/problems/runs/contests need time field.

		user 		-> 		register_time
		problem 	-> 		create_time
		contest 	->		create_time
		run 		-> 		submit_time

###user:

	login:
		request:
			POST /user/login
			(uid, pwd)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"uid":					<user_id>
					"ok":					<is_successful>
					"privilege":			<privilege>
					"status":				<account_status>
				}
			}
	logout:
		request:
			POST /user/logout
			(uid)
		response:
			{
				"head": 			<http_head>
				"status": 			<status_code>
				"error":			<error_description>
				"data":
				{
					"uid":			<user_id>
					"ok":			<is_successful>
				}
			}
	list:
		request: /user/list
		(uid, name, privilege, status, ORDER, OFFSET, LIMIT)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"list":
					[
						"uid":				<user_id>
						"name":				<nick_name>
						"privilege":		<privilege>
						"status":			<account_status>
					]
					...
				}
			}
	detail:
		request:
			POST /user/detail/uid/<uid>
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"uid" :					<user_id>
					"name":					<nick_name>
					"signature":			<user_description>
					"from""					<location_or_school>
					"email":				<email>
					"privilege":			<privilege>
					"status":				<account_status>
				}
			}
	insert:
		request:
			POST: /user/insert
			(uid, pwd, name, signature, from, email, status)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"uid":					<user_id>
					"ok":					<is_successful>
					"privilege":			<privilege>
					"status":				<account_status>
				}
			}
	edit:
		request:
			POST: /user/edit
			(uid, pwd, name, signature, from, email, status)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"uid":					<user_id>
					"ok":					<is_successful>
					"privilege":			<privilege>
					"status":				<account_status>
				}
			}
	delete:
		request:
			POST: /user/delete/uid/<uid>
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"uid":					<user_id>
					"ok":					<is_successful>
				}
			}

###problem:

	list:
		request:
			POST /problem/list
			(pid, *title*, *source*, status, ORDER, OFFSET, LIMIT)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"list":
					[
						"pid":				<problem_id>
						"title":			<title>
						"status":			<problem_status>	
					]
					...
				}
			}
	detail:
		request:
			POST /problem/detail/pid/<pid>
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"pid":					<problem_id>
					"time":					<time_limit>
					"memory":				<memory_limit>
					"title":				<title>
					"description":  		<description>
					"input":				<input>
					"output":				<output>
					"sample_input":			<sample_input>
					"sample_output":		<sample_output>
					"source":				<source>
					"hint":					<hint>
					"status":				<problem_status>
				}
			}
	insert:
		request:
			POST /problem/insert
			(time, memory, title, description, input, output, sample_input, sample_output, source, hint, status)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"pid":					<problem_id>
					"ok":					<is_successful>
					"status":				<problem_status>
				}
			}
	edit:
		request:
			POST /problem/edit/pid/<pid>
			(time, memory, title, description, input, output, sample_input, sample_output, source, hint, status)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"pid":					<problem_id>
					"ok":					<is_successful>
					"status":				<problem_status>
				}
			}
	delete:
		request:
			POST /problem/delete/pid/<pid>
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"pid":					<problem_id>
					"ok":					<is_successful>
				}
			}

###run:

	list:
		request:
			POST /run/list
			(rid, pid, uid, type, set, result, language, status, ORDER, OFFSET, LIMIT)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					list:
					[
						"rid":				<run_id>
						"pid":				<problem_id>
						"uid":				<user_id>
						"result":			<judge_result>
						"language":			<code_language>
						"time":				<run_time>
						"memory":			<run_memory>
						"length":			<code_length>
						"type":				<run_type>
						"set":				<belong_to>
						"status":			<run_status>
					]
					...
				}
			}
	insert:
		request:
			POST /run/insert
			(pid, uid, result, language, time, memory, length, tyoe, set, status)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"rid":					<run_id>
					"ok":					<is_successful>
					"status":				<run_status>
				}
			}
	edit:
		request:
			POST /run/edit/rid/<rid>
			(pid, uid, result, language, time, memory, length, tyoe, set, status)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"rid":					<run_id>
					"ok":					<is_successful>
					"status":				<run_status>
				}
			}
	delete:
		request:
			POST /run/delete/rid/<rid>
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"rid":					<run_id>
					"ok":					<is_successful>
				}
			}

###contest:

	list:
		request:
			POST /contest/list
			(cid, *title*, type, start, end, status, ORDER, OFFSET, LIMIT)
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					list:
					[
						"cid":				<contest_id>
						"title":			<title>
						"type":				<contest_type>
						"start":			<start_time>
						"end":				<end_time>
						"status":			<contest_status>
					]
					...
				}
			}
	detail:
		request:
			POST /contest/detail/cid/<cid>
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"cid":					<contest_id>
					"title":				<contest_title>
					"type":					<contest_type>
					"start":				<start_time>
					"end":					<end_time>
					"status":				<contest_status>
					list:
					[
						"pid":				<problem_id>
					]
					...
				}
			}
	insert:
		request:
			POST /contest/insert
			(title, type, start, end, status, list[pid...])
		response: 
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"cid":					<contest_id>
					"ok":					<is_successful>
					"status":				<contest_status>
				}
			}
	edit:
		request:
			POST /contest/edit/cid/<cid>
			(title, type, start, end, status, list[pid...])
		response: 
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"cid":					<contest_id>
					"ok":					<is_successful>
					"status":				<contest_status>
				}
			}
	delete:
		request:
			POST /contest/delete/cid/<cid>
		response:
			{
				"head": 					<http_head>
				"status": 					<status_code>
				"error":					<error_description>
				"data":
				{
					"cid":					<contest_id>
					"ok":					<is_successful>
				}
			}

###exercise:

	same as contest
