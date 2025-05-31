<?php

$command = "ping -c1 example.com";

$forbidden = array(
    'command',
    'output',
);

foreach ($_GET as $key => $value) {
    if (!in_array($key, $forbidden)) {
        $GLOBALS[$key] = $value;
    }
}

exec($command, $output);

echo "<p>" . htmlspecialchars($message) . "</p>";

foreach ($output as $line) {
    echo "<p>" . htmlspecialchars($line) . "</p>";
}

?>
