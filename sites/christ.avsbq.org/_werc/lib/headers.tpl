<!DOCTYPE HTML>
<html>
<head>
    <title>%($pageTitle%)</title>
    <meta charset="UTF-8">
% if(! ~ $#meta_description 0)
%   echo '    <meta name="description" content="'$"meta_description'">'
% if(! ~ $#meta_keywords 0)
%   echo '    <meta name="keywords" content="'$"meta_keywords'">'
% h = `{get_lib_file headers.inc}
% if(: ~ $#h 0)
%   cat $h
    %($"extraHeaders%)
</head>
<body>

