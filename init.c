#include <emacs-module.h>

int plugin_is_GPL_compatible;

extern emacs_value Fb64_encode(emacs_env *env, ptrdiff_t nargs, emacs_value args[], void *data);

extern emacs_value Fb64_decode(emacs_env *env, ptrdiff_t nargs, emacs_value args[], void *data);

static void
bind_function (emacs_env *env, const char *name, emacs_value Sfun)
{
        emacs_value Qfset = env->intern (env, "fset");
        emacs_value Qsym = env->intern (env, name);

        emacs_value args[] = { Qsym, Sfun };

        env->funcall (env, Qfset, 2, args);
}

static void
provide (emacs_env *env, const char *feature)
{
        emacs_value Qprovide = env->intern (env, "provide");
        emacs_value Qfeat = env->intern (env, feature);

        emacs_value args[] = { Qfeat };

        env->funcall (env, Qprovide, 1, args);
}

int
emacs_module_init (struct emacs_runtime *ert)
{
        emacs_env *env = ert->get_environment (ert);

#define DEFUN(lsym, csym, amin, amax, doc, data)                        \
        bind_function (env, lsym,                                       \
                       env->make_function(env, amin, amax, csym, doc, data))
        DEFUN ("b64-encode", Fb64_encode, 1, 1, "Return base64 encoded string.", NULL);
        DEFUN ("b64-decode", Fb64_decode, 1, 1, "Return base64 decoded string.", NULL);
#undef DEFUN
        
        provide (env, "go-base64");

        return 0;
}
