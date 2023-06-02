-- Get all targets for all topics
select tt.name, tt.consume_type, td.target from tb_dispatch td left join tb_topic tt on td.topic_id = tt.id;
