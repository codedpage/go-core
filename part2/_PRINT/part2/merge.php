<?php

	
	$path = "."; //home/user/public/foldername/";
	
	//Open the folder
	
	$dir_handle = @opendir($path) or die("Unable to open $path");
	
	//file & folders
	$files 		= array();
	$folders 	= array();


    // Loop through the files

    while ($file = readdir($dir_handle)) 
	{
	    if($file == "." || $file == ".." || $file == "merge.php"  || $file == "splitter.php"  || $file == "main")

        continue;
		
		if (is_file($file))
			$files[] = $file;
		else
			$folders[] = $file;
						
        //echo "<a href=\"$file\">$file</a><br />";
    }
	
	foreach($files as $f) 
	{
		#$re_arrange[filemtime($f)] = $f;
		$re_arrange[] = $f;
	}
	//krsort($re_arrange);
	
	foreach($re_arrange as $f) 
	{
		 
		echo $f; 
		echo "<br>";
		$content = "==".$f."========================\n".file_get_contents("$f")."\n";
		file_put_contents("main", $content, FILE_APPEND);		
		 
	}
    closedir($dir_handle);
	echo "<b>Done</b>";

?> 

 




