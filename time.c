/*
 * Time get the entire time execution of a c-file. It makes 
 * use of pipe(), fork(), execlp() and wait() System Calls.
 */
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/time.h>

#define PROGRAM_NAME  "time"

int
main (int argc, char *argv[])
{
  struct timeval start, end;
  char *argvx[argc - 1];
  int fd[2];

  if (argc == 1)
    {
      printf ("Usage: %s -C FILE [OPTION]\n", PROGRAM_NAME);
      exit (1);
    }
  if (pipe (fd) < 0)
    {
      fprintf (stderr, "%s: pipe() failed\n", PROGRAM_NAME);
      exit (1);
    }
  switch (fork ())
    {
    case -1:
      fprintf (stderr, "%s: fork() failed\n", PROGRAM_NAME);
      exit (1);
    case 0:
      close (fd[0]);
      for (int j = 0, i = 1; (argvx[j] = argv[i]) != NULL; i++, j++);
      /*
       * All statements after execlp call are ignored if it's 
       * executed successfully, that's why it's needed getting 
       * the time two instrucctions before execlp is called.
       */
      gettimeofday (&start, NULL);
      write (fd[1], &start, sizeof (struct timeval));
      execvp (argvx[0], argvx);
      fprintf (stderr, "%s: exec(\"%s\") failed\n", argvx[0], PROGRAM_NAME);
      exit (1);
    default:
      wait (NULL);
      gettimeofday (&end, NULL);
      close (fd[1]);
      if ((read (fd[0], &start, sizeof (struct timeval))) == -1)
	{
	  fprintf (stderr, "%s: could not read from child pipe\n",
		   PROGRAM_NAME);
	  exit (1);
	}
      printf ("\nElapsed %s: %f ms\n", PROGRAM_NAME,
	      (end.tv_sec - start.tv_sec) * 1000.0 + (end.tv_usec -
						      start.tv_usec) /
	      1000.0);
      close (fd[0]);
      exit (0);
    }
}
