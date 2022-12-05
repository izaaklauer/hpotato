# A sample local config file

server {
  bind_addr = ":8080"
}

%%wp_project%% {
  hello_world_message = "hello from %%wp_project%%"
}