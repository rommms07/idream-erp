<pre>
  [user]
===============================================
  id              |   uint64
  fb_id           |   uint64
  uname           |   string
  full_name       |   string
  email           |   string
  mobile          |   string
  birthdate       |   google.protobuf.Timestamp
  created_at      |   google.protobuf.Timestamp
  gender          |   enum_gender
  type            |   enum_type
==============================================

  [sessions]
==============================================
  user_id         |   uint64
  fb_access_token |   string
  fb_expires_in   |   string
  fb_token_type   |   enum token_type
  gen_uuid        |   string
  login_at        |   google.protobuf.Timestamp
===============================================

  [signin_logs]
=================================================
  user_id         |   uint64
  gen_uuid        |   string
  fb_token_type   |   string
  login_at        |   google.protobuf.Timestamp
=================================================

</pre>